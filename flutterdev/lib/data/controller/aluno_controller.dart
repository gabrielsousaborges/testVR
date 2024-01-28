import 'dart:developer';

import 'package:flutterdev/data/model/model_aluno.dart';
import 'package:flutterdev/data/repositories/aluno_repository.dart';

class AlunoController {
  final AlunoRepository alunoRepository;

  AlunoController(this.alunoRepository);

  bool isLoading = true;

  AlunoModel? _loadAluno;

  Future<void> carregaAluno(int id) async {
    try {
      final aluno = await alunoRepository.getAluno(id);

      _loadAluno = aluno;
    } catch (e) {
      log("Erro carrega Aluno");
    }
  }

  Future<void> cadastrarAluno(AlunoModel aluno) async {
    try {
      await alunoRepository.postAluno(aluno);
    } catch (e) {
      log('Erro ao cadastrar aluno: $e');
    }
  }
}
