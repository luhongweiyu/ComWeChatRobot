// ConsoleApplication1.cpp : 此文件包含 "main" 函数。程序执行将在此处开始并结束。
//

#include <iostream>
#include<windows.h>


//#pragma comment(lib, "wxDriver.lib")                                     // 告诉程序lib文件的路径，这里就表示当前目录
//extern "C" __declspec(dllimport) int __stdcall new_wechat();       // 这里使用__declspec(dllimport)，正好和生成dll时的__declspec(dllexport)对应
////extern "C" __declspec(dllimport) int __stdcall subtruct(int a, int b);  // 生成dll时没有加__stdcall的话，这里也不用加


int main()
{
	// 运行时加载DLL库
	HMODULE module = LoadLibrary(L"wxDriver.dll");     // 根据DLL文件名，加载DLL，返回一个模块句柄
	if (module == NULL)
	{
		printf("加载wxDriver.dll动态库失败\n");
		return 1;
	}
	typedef int(*Func_new_wechat)();                  // 定义函数指针类型
	Func_new_wechat new_wechat;
	// 导出函数地址
	new_wechat = (Func_new_wechat)GetProcAddress(module, "new_wechat");     // GetProcAddress返回指向的函数名的函数地址

	typedef int(*Func_start_listen)(int,int);
	Func_start_listen start_listen ;
	start_listen = (Func_start_listen)GetProcAddress(module, "start_listen");


	typedef int(*Func_stop_listen)(int);
	Func_stop_listen stop_listen;
	stop_listen = (Func_stop_listen)GetProcAddress(module, "stop_listen");


	//std::cout << "请输入端口,默认端口为:8686\n";

	int pid = new_wechat();
	int port = 8686;
    std::cout <<"PID:"<< pid <<"端口:"<< port << " !\n";

	std::cout << "开始运行\n";

	start_listen(pid, port);

	//stop_listen(pid);

	std::cout << "结束运行\n";
	getchar();
	//std::cout << pid << "结束运行 World!\n";

	return 1;
}

// 运行程序: Ctrl + F5 或调试 >“开始执行(不调试)”菜单
// 调试程序: F5 或调试 >“开始调试”菜单

// 入门使用技巧: 
//   1. 使用解决方案资源管理器窗口添加/管理文件
//   2. 使用团队资源管理器窗口连接到源代码管理
//   3. 使用输出窗口查看生成输出和其他消息
//   4. 使用错误列表窗口查看错误
//   5. 转到“项目”>“添加新项”以创建新的代码文件，或转到“项目”>“添加现有项”以将现有代码文件添加到项目
//   6. 将来，若要再次打开此项目，请转到“文件”>“打开”>“项目”并选择 .sln 文件
