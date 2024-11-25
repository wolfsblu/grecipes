import createClient from "openapi-fetch";
import type {paths} from "../../../api";

const client = createClient<paths>({baseUrl: "/api/"})

export const fetchProfile = async () => {
    return client.GET("/user/profile/")
}

export const fetchRecipes = async () => {
    return client.GET("/recipes")
}

export const createRecipe = (recipe: Recipe) => {
    return client.POST("/recipes", {
        body: {
            name: recipe.name,
            minutes: recipe.minutes,
            servings: recipe.servings,
        }
    })
}

export const login = (credentials: Credentials) => {
    return client.POST("/login", {
        body: {
            email: credentials.email,
            password: credentials.password,
        }
    })
}