class CalendsException implements Exception {
  CalendsException([String message = 'Unknown error']) {
    _message = message;
  }

  late String _message;

  @override
  String toString() {
    return _message;
  }
}
