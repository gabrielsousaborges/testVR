import 'dart:convert';
import 'dart:developer';

import 'package:flutterdev/data/model/model_aluno.dart';
import 'package:http/http.dart';

import '../../utils/exception.dart';
import '../repositories/aluno_repository.dart';

class Http implements AlunoRepository {
  final Client _client;

  Http({required Client client}) : _client = client;

  @override
  Future<AlunoModel?> getAluno(int id) async {
    try {
      final url = '$API_URL/$id';

      final response = await _client.get(Uri.parse(url));

      if (response.statusCode == 200) {
        return AlunoModel.fromJson(response.body);
      } else {
        throw ApiException(message: 'Erro buscar aluno');
      }
    } catch (e) {
      log("Erro ao buscar aluno");
      throw ApiException(message: 'Erro buscar aluno');
    }
  }

  @override
  Future<void> postAluno(AlunoModel aluno) async {
    try {
      final url = '$API_URL/alunos';

      final response = await _client.post(
        Uri.parse(url),
        body: json.encode(aluno.toJson()),
        headers: {'Content-Type': 'application/json'},
      );

      if (response.statusCode != 200) {
        throw ApiException(message: 'Erro ao adicionar aluno');
      }
    } catch (e) {
      log('Erro ao adicionar aluno: $e');
      throw ApiException(message: 'Erro ao adicionar aluno');
    }
  }
}
