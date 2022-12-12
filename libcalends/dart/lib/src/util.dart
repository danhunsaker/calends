import 'dart:ffi';
import 'dart:io';
import 'dart:typed_data';

import 'package:calends/src/calends_exception.dart';
import 'package:ffi/ffi.dart';
import 'package:path/path.dart' as path;

import 'generated_bindings.dart';

void panicHandler(Pointer<Char> message) {
  throw CalendsException(charPtrToStr(message));
}

CalendsBindings loadLib() {
  var libraryName = 'libcalends.so'; // Linux and Android
  if (Platform.isMacOS) libraryName = 'libcalends.dylib';
  if (Platform.isWindows) libraryName = 'libcalends.dll';

  final libraryPath = path.join(Directory.current.path, '..', libraryName);

  return CalendsBindings(
    Platform.isIOS
        ? DynamicLibrary.process()
        : DynamicLibrary.open(libraryPath),
  );
}

Pointer<Void> bytesToVoidPtr(Uint8List data) {
  final Pointer<Uint8> dataPtr = malloc<Uint8>(data.length);
  final pointerList = dataPtr.asTypedList(data.length);
  pointerList.setAll(0, data);

  return Pointer<Void>.fromAddress(dataPtr.address);
}

Uint8List voidPtrToBytes(Pointer<Void> data, int length) {
  final dataPtr = Pointer<Uint8>.fromAddress(data.address);
  return dataPtr.asTypedList(length);
}

Pointer<Char> strToCharPtr(String str) {
  return Pointer<Char>.fromAddress(str.toNativeUtf8().address);
}

String charPtrToStr(Pointer<Char> ptr) {
  if (ptr == nullptr) return '';
  return (Pointer<Utf8>.fromAddress(ptr.address)).toDartString();
}
