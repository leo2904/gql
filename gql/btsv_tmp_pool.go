// Code generated by " ../../../../github.com/grailbio/base/gtl/generate.py --PREFIX=btsvTmp --package=gql --output=btsv_tmp_pool.go -DELEM=btsvStructTmp -DMAXSIZE=128 ../../../../github.com/grailbio/base/gtl/freepool.go.tpl ". DO NOT EDIT.

package gql

// A freepool for a single thread. The interface is the same as sync.Pool, but
// it avoids locks and interface conversion overhead.
//
// Example:
//  generate.py -package=foo -prefix=int -DbtsvStructTmp=foo -D128=128
//
//
// Parameters:
//  btsvStructTmp: the object to be kept in the freepool
//  128: the maxium number of objects to keep in the freepool

type btsvTmpPool struct {
	New func() btsvStructTmp
	p   []btsvStructTmp
}

func (p *btsvTmpPool) Get() btsvStructTmp {
	if len(p.p) == 0 {
		return p.New()
	}
	tmp := p.p[len(p.p)-1]
	p.p = p.p[:len(p.p)-1]
	return tmp
}

func (p *btsvTmpPool) Put(tmp btsvStructTmp) {
	if len(p.p) < 128 {
		p.p = append(p.p, tmp)
	}
}
