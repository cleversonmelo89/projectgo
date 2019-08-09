# Projectgo Gestão de Tags Repositório Git

Para este  projeto é necessário utilizar o banco de dados NoSQL MongoDB para inclusão, alteração, consulta e exclusão  dos dados.

<h1>Versões utilizadas</h1>

-go1.12.7<br>
-Banco de dados MongoDB v4.0.10<br>
-IDE JetBrains GoLand 2019.2<br>

<h1> EndPoints Disponiveis </h1>

<h2>Recuperar repositórios GitHub por usuário</h2>

<h2>GET</h2>
http://localhost:3000/api/v1/repo/git/{user}
<p>Response Body</p>

```json
[
    {
        "bson_id": "",
        "id": 138361864,
        "html_url": "https://github.com/gabrielpborba/SubAndroidApi",
        "name": "SubAndroidApi",
        "description": "Trab Sub de Android",
        "language": "Java",
        "tags": null
    }
]
```

<h2>Recuperar todos os repositórios cadastrado na base local</h2>

<h2>GET</h2>
http://localhost:3000/api/v1/repo
<p>Response Body</p>

```json
{
    "Total": 2,
    "Repo": [
        {
            "bson_id": "5d4b28f0485d181250000bdd",
            "id": 72203974,
            "html_url": "https://github.com/sqreen/awesome-nodejs-projects",
            "name": "awesome-nodejs-projects",
            "description": "Curated list of awesome open-source applications made with Node.js",
            "language": "",
            "tags": [
                {
                    "tag_name": "spring boot"
                },
                {
                    "tag_name": "spring core"
                }
            ]
        },
        {
            "bson_id": "5d4b28f0485d181250000bde",
            "id": 574877,
            "html_url": "https://github.com/aws/aws-sdk-java",
            "name": "aws-sdk-java",
            "description": "The official AWS SDK for Java.",
            "language": "Java",
            "tags": []
        }
    ]
 }
```

<h2>Recuperar repositórios por tags</h2>

<h2>GET</h2>
http://localhost:3000/api/v1/repo/tag
<p>Request Body</p>

```json
[
  {
    "tag_name": "core"
  },
  {
    "tag_name": "boot"
  }
]
```

<p>Response Body</p>

```json
[
    {
        "bson_id": "5d4b28f0485d181250000bdd",
        "id": 72203974,
        "html_url": "https://github.com/sqreen/awesome-nodejs-projects",
        "name": "awesome-nodejs-projects",
        "description": "Curated list of awesome open-source applications made with Node.js",
        "language": "",
        "tags": [
            {
                "tag_name": "spring boot"
            },
            {
                "tag_name": "spring core"
            }
        ]
    }
]
```

<h2>Recuperar repositórios pelo id do Git</h2>

<h2>GET</h2>
http://localhost:3000/api/v1/repo/{id_git}
<p>Response Body</p>

```json
[
    {
        "bson_id": "5d4b28f0485d181250000bdd",
        "id": 72203974,
        "html_url": "https://github.com/sqreen/awesome-nodejs-projects",
        "name": "awesome-nodejs-projects",
        "description": "Curated list of awesome open-source applications made with Node.js",
        "language": "",
        "tags": [
            {
                "tag_name": "spring boot"
            },
            {
                "tag_name": "spring core"
            }
        ]
    }
]
```

<h2>Adicionar tags ao repositório</h2>

<h2>POST</h2>
http://localhost:3000/api/v1/repo/{id_git}/addTag
<p>Request Body</p>

```json
[
  {
    "tag_name": "backend"
  },
  {
    "tag_name": "spring boot"
  }
]
```

<h2>Deletar tags do repositório</h2>

<h2>DELETE</h2>
http://localhost:3000/api/v1/repo/{id_git}/deleteTag
<p>Request Body</p>

```json
[
  {
    "tag_name": "backend"
  },
  {
    "tag_name": "spring boot"
  }
]
```

<h2>Editar uma tag de um repositório</h2>

<h2>PATCH</h2>
http://localhost:3000/api/v1/repo/{id_git}/editTag/tag/{tag_name}/new/{tag}"

<h2>Obter sugestões de tags de um repositório</h2>
<h2>GET</h2>
http://localhost:3000/api/v1/repo/tag/suggestions/{id_git}

```json
[
    {
        "suggestion_name": "awesome"
    },
    {
        "suggestion_name": "nodejs"
    },
    {
        "suggestion_name": "projects"
    }
]
```