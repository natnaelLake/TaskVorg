package gql

const CreateTask = `
mutation CreateTask($title: String!, $description: String, $userId: Int!, $status: String) {
    insert_tasks(objects: {title: $title, description: $description, user_id: $userId, status: $status}) {
        returning {
            id
            title
            description
            status
            created_at
            updated_at
        }
    }
}
`

const CreateUsers = `
mutation CreateUser($name: String!, $email: String!) {
    insert_users(objects: {name: $name, email: $email}) {
        returning {
            id
            name
            email
        }
    }
}
`


