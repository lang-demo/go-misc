#include "goDLL.h"
#include <string>
#include <iostream>
#include "exportgo.h"
#include "strconv2.h"

// force gcc to link in go runtime (may be a better solution than this)
void dummy() {
	__ThisPath__();
}

GoString to_gostr(const std::string &s)
{
  GoString gs;
  gs.p = s.c_str();
  gs.n = s.size();
  return gs;
}

std::string from_gostr(const GoString &s)
{
  std::string result(s.p, s.n);
  return result;
}

int main() {
  unicode_ostream uout(std::cout, GetConsoleOutputCP());

  printf("GetIntFromDLL(): %d\n", GetIntFromDLL());
  PrintBye();
  std::string name = "xyz helloハロー©";
  PrintHello(to_gostr(name));
  GoString result = GetStringFromDLL();
  uout << "from_gostr(result): " << from_gostr(result) << std::endl;
  GoUintptr uptr = GetStringPtrFromDLL();
  uout << "GetStringPtrFromDLL(): " << (const wchar_t *)uptr << std::endl;
  auto uptr2 = GetStringPtrFromDLL2();
  uout << "GetStringPtrFromDLL2(): " << (const char *)uptr2 << std::endl;
  PrintHelloPtr((GoUintptr)L"トム©");
  PrintHelloPtr2((GoUintptr)u8"マリー©");

  uout << "ThisPath() " << (const char *)(__ThisPath__()) << std::endl;

  return 0;
}
