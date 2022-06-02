package auth

import (
	"errors"
	"github.com/go-ldap/ldap/v3"
	"log"
	"sync"
)

var (
	InvalidCapacityError = errors.New("invalid capacity settings")
	ErrPoolClosed        = errors.New("pool closed")
	ManagerDNBindError   = errors.New("manager dn bind error")
)

type Pool struct {
	m       sync.Mutex                 // 保证多个goroutine访问时候，closed的线程安全
	res     chan ldap.Conn             // 连接存储的chan
	factory func() (*ldap.Conn, error) // 新建连接的工厂方法
	closed  bool                       // 连接池关闭标志
}

func NewPool(fn func() (*ldap.Conn, error), cap int) (*Pool, error) {
	if cap <= 0 {
		return nil, InvalidCapacityError
	}
	return &Pool{
		res:     make(chan ldap.Conn, cap),
		factory: fn,
		closed:  false,
	}, nil
}

func (p *Pool) Acquire() (*ldap.Conn, error) {
	select {
	case r, ok := <-p.res:
		log.Println("Acquire:共享资源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return &r, nil
	default:
		log.Println("Acquire:新生成资源")
		return p.factory()
	}
}

// Close 关闭资源池，释放资源
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	//关闭通道，不让写入了
	close(p.res)
	//关闭通道里的资源
	for r := range p.res {
		r.Close()
	}
}

func (p *Pool) Release(c ldap.Conn) {
	//保证该操作和Close方法的操作是安全的
	p.m.Lock()
	defer p.m.Unlock()

	//资源池都关闭了，就省这一个没有释放的资源了，释放即可
	if p.closed {
		c.Close()
		return
	}
	select {
	case p.res <- c:
		log.Println("资源释放到池子里了")
	default:
		log.Println("资源池满了，释放这个资源吧")
		c.Close()
	}
}
