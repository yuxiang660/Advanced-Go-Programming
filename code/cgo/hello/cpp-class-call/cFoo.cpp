#include "cFoo.h"
#include "Foo.h"

cFoo openCFoo()
{
    Foo* ret = new Foo();
    return (cFoo)(ret);
}

void closeCFoo(cFoo cf)
{
    Foo* f = (Foo*)cf;
    delete f;
}

void cfoo(cFoo cf)
{
    Foo* f = (Foo*)cf;
    f->foo();
}
