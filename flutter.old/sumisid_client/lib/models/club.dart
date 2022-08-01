class Club {
  final int id;
  final String name;
  final String streetAddress;
  final String streetNumber;
  final String zipCode;
  final String phoneNumber;
  final int contactPersonId;
  final String extraInfo;

  const Club(
      {required this.id,
      required this.name,
      required this.streetAddress,
      required this.streetNumber,
      required this.zipCode,
      required this.phoneNumber,
      required this.contactPersonId,
      required this.extraInfo});

  factory Club.fromJson(Map<String, dynamic> json) {
    return Club(
      id: json['id'],
      name: json['name'],
      streetAddress: json['streetAddress'],
      streetNumber: json['streetNumber'],
      zipCode: json['zipCode'],
      phoneNumber: json['phoneNumber'],
      contactPersonId: json['contactPersonId'],
      extraInfo: json['extraInfo'],
    );
  }
}
