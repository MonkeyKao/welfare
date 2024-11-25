import 'package:flutter/material.dart';
import 'package:mobile_scanner/mobile_scanner.dart';

class BarcodeScannerSimple extends StatefulWidget {
  const BarcodeScannerSimple({super.key});

  @override
  State<BarcodeScannerSimple> createState() => _BarcodeScannerSimpleState();
}

class _BarcodeScannerSimpleState extends State<BarcodeScannerSimple> {
  Barcode? _barcode; // 用於顯示最近掃描結果
  bool _hasScanned = false; // 避免重複處理掃描結果

  Widget _buildBarcode(Barcode? value) {
    if (value == null) {
      return const Text(
        '請掃描QR Code！',
        overflow: TextOverflow.fade,
        style: TextStyle(color: Colors.white),
      );
    }

    return Text(
      value.displayValue ?? '未獲取有效值。',
      overflow: TextOverflow.fade,
      style: const TextStyle(color: Colors.white),
    );
  }

  void _handleBarcode(BarcodeCapture barcodes) {
    if (_hasScanned) return; // 如果已經處理過掃描結果則忽略

    final Barcode? barcode = barcodes.barcodes.firstOrNull;
    if (barcode != null) {
      setState(() {
        _barcode = barcode; // 更新界面上顯示的條形碼信息
        _hasScanned = true; // 標記為已處理
      });

      // 返回掃描結果到上一頁
      Navigator.pop(context, barcode.displayValue);

      // 可選：可以顯示一個提示用戶信息掃描完成的彈窗
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('掃描成功：${barcode.displayValue}')),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('請掃描QR Code')),
      backgroundColor: Colors.black,
      body: Stack(
        children: [
          MobileScanner(
            onDetect: _handleBarcode,
          ),
          Align(
            alignment: Alignment.bottomCenter,
            child: Container(
              alignment: Alignment.bottomCenter,
              height: 100,
              color: Colors.black.withOpacity(0.4),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                children: [
                  Expanded(child: Center(child: _buildBarcode(_barcode))),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}