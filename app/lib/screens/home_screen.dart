import 'package:flutter/material.dart';
import 'package:webview_flutter/webview_flutter.dart';
import 'package:logger/logger.dart';

final logger = Logger();

class HomeScreen extends StatefulWidget {
  const HomeScreen({super.key});

  @override
  State<StatefulWidget> createState() => _HomeScreen();
}

class _HomeScreen extends State<HomeScreen> {
  late final WebViewController controller;
  final String initialUrl = 'http://192.168.0.239:5000'; // 初始 URL
  bool showBackButton = false; // 是否顯示返回按鈕

  @override
  void initState() {
    super.initState();
    controller = WebViewController()
      ..setJavaScriptMode(JavaScriptMode.unrestricted)
      ..setNavigationDelegate(
        NavigationDelegate(
          onPageStarted: (url) {
            logger.d('Page started loading: $url');
            _updateBackButtonVisibility(url);
          },
          onPageFinished: (url) {
            logger.d('Page finished loading: $url');
            _updateBackButtonVisibility(url);
          },
        ),
      )
      ..loadRequest(Uri.parse(initialUrl));
  }

  // 更新返回按鈕的顯示狀態
  void _updateBackButtonVisibility(String url) {
    setState(() {
      // 當 URL 包含 initialUrl 時，隱藏返回按鈕，否則顯示
      showBackButton = !url.contains(initialUrl);
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: showBackButton
          ? AppBar(
              leading: IconButton(
                icon: const Icon(Icons.arrow_back),
                onPressed: () async {
                  // 返回初始頁面
                  await controller.loadRequest(Uri.parse(initialUrl));
                  setState(() {
                    showBackButton = false; // 返回初始頁面後隱藏返回按鈕
                  });
                },
              ),
            )
          : null,
      body: WebViewWidget(controller: controller),
    );
  }
}