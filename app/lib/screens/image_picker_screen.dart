// import 'dart:io';

// import 'package:flutter/material.dart';
// import 'package:image_picker/image_picker.dart';

// class ImagePickerScreen extends StatefulWidget {
//   const ImagePickerScreen({super.key});

//   @override
//   State<StatefulWidget> createState() => _ImagePickerScreen();
// }

// class _ImagePickerScreen extends State<ImagePickerScreen> {
//   File? _selectedImage; // 存储用户选择的图片
//   final ImagePicker _picker = ImagePicker();

//   @override
//   Widget build(BuildContext context) {
//     return Scaffold(
//       appBar: AppBar(
//         title: const Text('选择图片'),
//         backgroundColor: Colors.black,
//       ),
//       body: Stack(
//         children: [
//           // 背景显示选择的图片
//           Positioned.fill(
//             child: _selectedImage == null
//                 ? const Center(
//                     child: Text(
//                       '未选择图片',
//                       style: TextStyle(color: Colors.grey, fontSize: 18),
//                     ),
//                   )
//                 : _buildImagePreview(),
//           ),
//           // 底部操作栏
//           Align(
//             alignment: Alignment.bottomCenter,
//             child: _buildBottomBar(),
//           ),
//         ],
//       ),
//     );
//   }
// }
