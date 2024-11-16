import router from "page"
import Index from "../pages/Index.svelte"
import About from "../pages/About.svelte"
import NotFound from "../pages/404.svelte"
import CreateRecipe from "../pages/recipes/Create.svelte"

let page = $state(Index);

export const createRouter = () => {

    const registerRoutes = () => {
        router("/", () => page = Index)
        router("/about", () => page = About)
        router("/recipes/create", () => page = CreateRecipe)
        router("*", () => page = NotFound)

        router.start()
    }

    return {
        get page() {
            return page
        },
        registerRoutes,
    }
}