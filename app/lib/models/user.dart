// lib/models/user.dart
class User {
  final int id;
  final String account;
  final String password;
  final String name;
  final String salt;
  final DateTime birthday;
  final int female; // 可以用 0 或 1 來表示性別
  final int location;
  final String email;

  User({
    required this.id,
    required this.account,
    required this.password,
    required this.name,
    required this.salt,
    required this.birthday,
    required this.female,
    required this.location,
    required this.email,
  });

  // 從 JSON 創建 User 對象的工廠方法
  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json['_id'] as int,
      account: json['account'] as String,
      password: json['password'] as String,
      name: json['name'] as String,
      salt: json['salt'] as String,
      birthday: DateTime.parse(json['birthday'] as String),
      female: json['female'] as int,
      location: json['location'] as int,
      email: json['email'] as String,
    );
  }

  // 將 User 對象轉換為 JSON 格式
  Map<String, dynamic> toJson() {
    return {
      '_id': id,
      'account': account,
      'password': password,
      'name': name,
      'salt': salt,
      'birthday': birthday.toIso8601String(), // 轉換為 ISO 8601 字符串
      'female': female,
      'location': location,
      'email': email,
    };
  }
}