interface Credentials {
    email: string
    password: string
}

interface Error {
    code: number
    message: string
}

interface User {
    id: number
    email: string
}

interface Recipe {
    name: string
    servings: number | null
    minutes: number | null
}