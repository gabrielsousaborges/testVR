import 'package:flutter/material.dart';
import 'package:flutterdev/screen/alunoscreen.dart';
import 'package:flutterdev/screen/cadastroscreen.dart';
import 'package:flutterdev/screen/cursoscreen.dart';

class Home extends StatefulWidget {
  const Home({super.key});

  @override
  State<Home> createState() => _HomeState();
}

class _HomeState extends State<Home> {
  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        appBar: AppBar(
          title: const Text('FlutterDEV'),
        ),
        drawer: Drawer(
          child: ListView(
            children: <Widget>[
              const DrawerHeader(
                decoration: BoxDecoration(color: Colors.orange),
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Text('Teste VR'),
                  ],
                ),
              ),
              Container(
                padding:
                    const EdgeInsets.symmetric(vertical: 5, horizontal: 20),
                child: GestureDetector(
                  child: const Row(
                    children: [
                      Icon(Icons.person_2_outlined),
                      SizedBox(
                        width: 5,
                      ),
                      Text('Alunos'),
                    ],
                  ),
                  onTap: () {
                    Navigator.pop(context);
                    Navigator.push(context,
                        MaterialPageRoute(builder: (context) => AlunoScreen()));
                  },
                ),
              ),
              const Divider(),
              Container(
                padding:
                    const EdgeInsets.symmetric(vertical: 5, horizontal: 20),
                child: GestureDetector(
                  child: const Row(
                    children: [
                      Icon(Icons.collections_bookmark_outlined),
                      SizedBox(
                        width: 5,
                      ),
                      Text('Cursos'),
                    ],
                  ),
                  onTap: () {
                    Navigator.pop(context);
                    Navigator.push(context,
                        MaterialPageRoute(builder: (context) => CursoScreen()));
                  },
                ),
              ),
              const Divider(),
              Container(
                padding:
                    const EdgeInsets.symmetric(vertical: 5, horizontal: 20),
                child: GestureDetector(
                  child: const Row(
                    children: [
                      Icon(Icons.account_box_outlined),
                      SizedBox(
                        width: 5,
                      ),
                      Text('Cadastro'),
                    ],
                  ),
                  onTap: () {
                    Navigator.pop(context);
                    Navigator.push(
                        context,
                        MaterialPageRoute(
                            builder: (context) => CadastroScreen()));
                  },
                ),
              )
            ],
          ),
        ),
        body: Container(
          alignment: Alignment.center,
          child: Padding(
            padding: EdgeInsets.all(50),
            child: Container(
              child: Text(
                'Teste VR',
                style: TextStyle(color: Colors.black, fontSize: 40),
              ),
            ),
          ),
        ),
      ),
    );
  }
}
