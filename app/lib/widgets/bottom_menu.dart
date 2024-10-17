// ignore: file_names
import 'package:flutter/material.dart';
import 'package:webview_flutter/webview_flutter.dart';

class BottomMneu extends StatefulWidget {
  const BottomMneu({super.key});

  @override
  State<StatefulWidget> createState() => _BottomMneu();

}

class _BottomMneu extends State<BottomMneu> {

  WebViewController controller = WebViewController()
  ..setJavaScriptMode(JavaScriptMode.unrestricted)
  ..loadRequest(Uri.parse('https://www.youtube.com/'));

  int currentIndex = 0;
  NavigationDestinationLabelBehavior labelBehavior =
      NavigationDestinationLabelBehavior.onlyShowSelected;
  

  @override
  Widget build(BuildContext context) {
    final themeData = Theme.of(context);
    final screenHeight = MediaQuery.of(context).size.height;
    return Scaffold(
      bottomNavigationBar: NavigationBar(
        labelBehavior: labelBehavior,
        onDestinationSelected: (int index) {
          setState(() {
            currentIndex = index;
          });
        },
        height: screenHeight*0.1,
        indicatorColor: Colors.amber,
        selectedIndex: currentIndex,
        destinations: const <Widget>[
          NavigationDestination(
            selectedIcon: Icon(Icons.home),
            icon: Icon(Icons.home_outlined),
            label: '首頁',
          ),
          NavigationDestination(
            icon: Icon(Icons.account_circle_outlined),
            selectedIcon: Icon(Icons.account_circle),
            label: '用戶',
          ),
          NavigationDestination(
            icon: Icon(Icons.question_answer_outlined),
            selectedIcon: Icon(Icons.question_answer),
            label: 'Q&A',
          ),
        ],
      ),
      body: <Widget>[
        Card(
          shadowColor: Colors.transparent,
          margin: const EdgeInsets.all(8.0),
          child: SizedBox.expand(
            child: Center(
              child: Text(
                'Home page',
                style: themeData.textTheme.titleLarge,
              ),
            ),
          ),
        ),
        WebViewWidget(controller: controller),
        Card(
          shadowColor: Colors.transparent,
          margin: const EdgeInsets.all(8.0),
          child: SizedBox.expand(
            child: Center(
              child: Text(
                'QA page',
                style: themeData.textTheme.titleLarge,
              ),
            ),
          ),
        ),
      ][currentIndex],
    );
  }
}
