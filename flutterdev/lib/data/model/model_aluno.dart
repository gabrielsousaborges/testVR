import 'dart:convert';

class AlunoModel {
  final int id;
  final String nome;

  AlunoModel({required this.id, required this.nome});

  factory AlunoModel.fromMap(Map<String, dynamic> map) {
    return AlunoModel(id: map['id'] ?? 0, nome: map['nome'] ?? '');
  }

  factory AlunoModel.fromJson(String source) =>
      AlunoModel.fromMap(json.decode(source));

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'nome': nome,
    };
  }
}
