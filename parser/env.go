package simple_parser

type Env struct {
	table map[string]*Symbol
	prev  *Env //形成向外查找变量的链表
}

func NewEnv(p *Env) *Env {
	return &Env{
		table: make(map[string]*Symbol),
		prev:  p,
	}
}

func (e *Env) Put(s string, sym *Symbol) {
	e.table[s] = sym
}

func (e *Env) Get(s string) *Symbol {
	//查询变量符号时，若当前表内未定义，想上一层查找
	for env := e; env != nil; env = e.prev {
		found, ok := env.table[s]
		if ok {
			return found
		}
	}

	return nil
}
