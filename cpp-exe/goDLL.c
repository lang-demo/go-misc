#include "goDLL.h"
#include "exportgo.h"

// force gcc to link in go runtime (may be a better solution than this)
void dummy() {
	__ThisPath__();
}
