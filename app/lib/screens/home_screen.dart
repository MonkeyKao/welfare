import 'package:flutter/material.dart';
import '../widgets/bottom_menu.dart';


class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      bottomNavigationBar: const BottomMneu(),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            const Text(
              'Welcome to the Home Screen!',
              style: TextStyle(fontSize: 24),
            ),
            const SizedBox(height: 20),
            ElevatedButton(
              onPressed: () {
                _showMessage(context);
              },
              child: const Text('Press Me'),
            ),
          ],
        ),
      ),
    );
  }

  void _showMessage(BuildContext context) {
    const snackBar =  SnackBar(content: Text('Button Pressed!'));
    ScaffoldMessenger.of(context).showSnackBar(snackBar);
  }
}