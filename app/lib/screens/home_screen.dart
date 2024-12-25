import 'dart:convert';

import 'package:app/screens/qr_scanner_screen.dart';
import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';
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
  final String initialUrl = 'http://172.16.72.83:5000/'; // 初始 URL
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
      ..addJavaScriptChannel(
        'FlutterChannel', // 添加 JavaScript 通道
        onMessageReceived: (message) {
          logger.d('Received message from Vue: ${message.message}');
          logger.d('接收開啟照片選擇器');
          if (message.message == 'openCamera') {
            _openQRCodeScanner(); // 根據消息開啟二維碼掃描功能
          }else if(message.message == 'openImagePicker') {
            logger.d('接收開啟照片選擇器');
            _selectedImage();
          }
        },
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

  // 打開二維碼掃描
  Future<void> _openQRCodeScanner() async {
    final result = await Navigator.push(context, MaterialPageRoute(builder: (context) => const BarcodeScannerSimple() ) );

    if(result!=null && result is String) {
      await controller.runJavaScript("window.receiveQRCodeResult('$result');");
      logger.d('QR Code Scanned: $result');
    }else {
      logger.d("QR Code Scnaning failed");
    }
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

  Future<void> _selectedImage() async {
  final ImagePicker picker = ImagePicker();
  final XFile? pickedFile = await picker.pickImage(source: ImageSource.gallery);

  if (pickedFile != null) {
    // 将图片转为 Base64 编码
    final bytes = await pickedFile.readAsBytes();
    final base64Image = base64Encode(bytes);

    // 调用 WebView 的 JavaScript 方法，将 Base64 图片数据传递给 Vue
    final script = "window.receiveImage('$base64Image');";
    await controller.runJavaScript(script);

    logger.d('Image selected and sent to Vue: $base64Image');
  } else {
    logger.d('Image selection cancelled');
  }
}
}