import 'package:flutter/material.dart';
import 'package:flutterdev/screen/alunoscreen.dart';
import 'package:flutterdev/screen/homescreen.dart';
import 'package:provider/provider.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return Provider<AlunoScreen>(
      create: (_) => const AlunoScreen(),
      child: MaterialApp(
        title: 'FlutterDev',
        debugShowCheckedModeBanner: false,
        theme: ThemeData(
          appBarTheme: AppBarTheme(color: Colors.orangeAccent),
          primaryColor: Colors.orange,
          primarySwatch: Colors.orange,
        ),
        home: const Home(),
      ),
    );
  }
}
