import 'package:flutter/material.dart';
import 'screens/home_screen.dart';  // 引入 HomeScreen

void main() {
  runApp( const SafeArea(child: MyApp()) );
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: '眸福利',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: const HomeScreen(),  // 設定 HomeScreen 為主頁
    );
  }
}