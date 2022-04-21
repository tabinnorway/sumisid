import 'club.dart';

class Person {
  final int id;
  final String email;
  final String firstName;
  final String middleName;
  final String lastName;
  final DateTime birthDate;
  final bool isAdmin;
  final String phoneNumber;
  final int mainClubId;
  final Club? mainClub;

  const Person(
      {required this.id,
      required this.email,
      required this.firstName,
      required this.middleName,
      required this.lastName,
      required this.birthDate,
      required this.isAdmin,
      required this.phoneNumber,
      required this.mainClubId,
      this.mainClub});

  factory Person.fromJson(Map<String, dynamic> json) {
    return Person(
      id: json['id'],
      email: json['email'],
      firstName: json['firstName'],
      middleName: json['middleName'],
      lastName: json['lastName'],
      birthDate: DateTime.parse(json['birthDate']),
      isAdmin: json['isAdmin'],
      phoneNumber: json['phoneNumber'],
      mainClubId: json['mainClubId'],
      mainClub: Club.fromJson(json['mainClub']),
    );
  }
}
