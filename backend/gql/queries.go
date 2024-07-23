package gql

const (
    // Define your GraphQL queries and mutations here
    createTaskMutation = `
        mutation($title: String!, $description: String, $status: String!, $userId: String!) {
            insert_tasks(objects: {title: $title, description: $description, status: $status, userId: $userId}) {
                returning {
                    id
                    title
                    description
                    status
                    created_at
                    updated_at
                    userId
                }
            }
        }
    `

    getTasksQuery = `
        query {
            tasks(order_by: {created_at: desc}) {
                id
                title
                description
                status
                created_at
                updated_at
                userId
            }
        }
    `

    updateTaskMutation = `
        mutation($id: ID!, $title: String!, $description: String, $status: String!) {
            update_tasks(where: {id: {_eq: $id}}, _set: {title: $title, description: $description, status: $status}) {
                returning {
                    id
                    title
                    description
                    status
                    created_at
                    updated_at
                    userId
                }
            }
        }
    `

    deleteTaskMutation = `
        mutation($id: ID!) {
            delete_tasks(where: {id: {_eq: $id}}) {
                returning {
                    id
                    title
                    description
                    status
                    created_at
                    updated_at
                    userId
                }
            }
        }
    `
)
