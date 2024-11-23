import createClient from "openapi-fetch";
import type {paths} from "../../../api";

const client = createClient<paths>({baseUrl: "/api/"})

export interface Recipe {
    name: string
    servings: number | null
    minutes: number | null
}

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