import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:http/http.dart' as http;

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Проверить пропуск',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: const MyHomePage(title: 'Проверка пропусков'),
      debugShowCheckedModeBanner: false,
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({Key? key, required this.title}) : super(key: key);

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int _counter = 0;

  void _incrementCounter() async {
    print("Hello");

    var request = http.Request('GET', Uri.parse('http://192.168.91.249:8088/'));

    http.StreamedResponse response = await request.send();

    if (response.statusCode == 200) {
      List<dynamic> user = jsonDecode(await response.stream.bytesToString());
    } else {
      print(response.reasonPhrase);
    }

    // setState(() {
    //   _counter++;
    // });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: Center(
        child: Column(
          children: [
            const SizedBox(
              height: 30,
            ),
            Container(
              decoration: BoxDecoration(
                // И тут указываем размер скругления границ
                // В данном случае скругление со всех сторон одинаковое
                border: Border.all(
                  width: 1,
                  color: Colors.blue,
                ),
                borderRadius: BorderRadius.circular(10),
              ),
              child: SizedBox(
                  width: 130,
                  height: 50,
                  child: Row(
                    children: <Widget>[
                      SizedBox(
                        width: 90,
                        height: 50,
                        child: Column(children: [
                          TextField(
                              textAlignVertical: TextAlignVertical.top,
                              style: const TextStyle(
                                  fontSize: 22.0,
                                  height: 2.0,
                                  fontWeight: FontWeight.bold,
                                  color: Colors.black),
                              textAlign: TextAlign.right,
                              decoration:
                                  const InputDecoration.collapsed(hintText: ""),
                              keyboardType: TextInputType.text,
                              inputFormatters: [
                                LengthLimitingTextInputFormatter(6),
                              ]),
                        ]),
                      ),
                      SizedBox(
                        width: 30,
                        height: 50,
                        child: Column(
                            mainAxisAlignment: MainAxisAlignment.start,
                            children: <Widget>[
                              const SizedBox(
                                height: 4,
                              ),
                              TextField(
                                  textAlign: TextAlign.center,
                                  decoration: const InputDecoration.collapsed(
                                      hintText: ""),
                                  keyboardType: TextInputType.number,
                                  style: const TextStyle(
                                      fontWeight: FontWeight.bold,
                                      color: Colors.black),
                                  inputFormatters: [
                                    LengthLimitingTextInputFormatter(2),
                                  ]),
                              Image.asset(
                                'icons/ru.png',
                                width: 20,
                              ),
                            ]),
                      ),
                    ],
                  )),
            ),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton.extended(
          onPressed: _incrementCounter, label: const Text("Проверить пропуск")),
    );
  }
}
