import 'package:flutterdev/data/model/model_aluno.dart';

const String API_URL = 'localhost';

abstract class AlunoRepository {
  Future<AlunoModel?> getAluno(int id);

  Future<void> postAluno(AlunoModel aluno);
}
