#pragma once

#ifdef __cplusplus
extern "C" {
#endif

    typedef void* cFoo;
    cFoo openCFoo();
    void closeCFoo(cFoo cf);
	void cfoo(cFoo cf);

#ifdef __cplusplus
}
#endif
