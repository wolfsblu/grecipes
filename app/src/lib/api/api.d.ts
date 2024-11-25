
interface Error {
    code: number
    message: string
}

interface Recipe {
    name: string
    servings: number | null
    minutes: number | null
}

interface Credentials {
    email: string
    password: string
}