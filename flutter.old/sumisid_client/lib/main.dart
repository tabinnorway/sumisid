import 'package:flutter/material.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:sumisid/models/club.dart';

import 'models/person.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Sumisid',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        primarySwatch: Colors.deepPurple,
      ),
      home: const MyHomePage(title: 'Sumisid'),
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

  void _incrementCounter() {
    setState(() {
      _counter++;
    });
  }

  Future<List<Club>> fetchDiveClubs() async {
    var jsonString = (await http.get(Uri.parse('http://localhost:8080/api/v1/clubs'))).body;
    var jsonList = jsonDecode(jsonString) as List;
    return jsonList.map((jsonList) => Club.fromJson(jsonList)).toList();
  }

  Future<List<Person>> fetchPeople() async {
    var jsonString = (await http.get(Uri.parse('http://localhost:8080/api/v1/people'))).body;
    var jsonList = jsonDecode(jsonString) as List;
    return jsonList.map((jsonList) => Person.fromJson(jsonList)).toList();
  }

  void _fabPushed() async {
    _incrementCounter();
    for (var club in await fetchDiveClubs()) {
      // ignore: avoid_print
      print('===> Found ${club.name}');
    }
    for (var person in await fetchPeople()) {
      // ignore: avoid_print
      print('===> Found ${person.firstName} ${person.lastName} - ${person.mainClub?.name ?? 'no club'}');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(
          widget.title,
          style: const TextStyle(fontFamily: 'Orbitron', fontWeight: FontWeight.w900, fontSize: 32),
        ),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text('You have pushed the button $_counter times'),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(onPressed: _fabPushed, tooltip: 'Increment', child: const Icon(Icons.add)),
    );
  }
}
