import router, {type Callback} from "page"
import {fetchProfile} from "./api/client";
import {type Component} from "svelte"
import About from "../pages/About.svelte"
import CreateRecipe from "../pages/recipes/Create.svelte"
import Index from "../pages/Index.svelte"
import Login from "../pages/Login.svelte";
import NotFound from "../pages/errors/404.svelte"
import Register from "../pages/Register.svelte";

let page: Component | null = $state(null);

export const createRouter = () => {
    const requireLogin: Callback = async (_, next) => {
        const redirectToLogin = () => router.redirect("/login")
        try {
            const profile = await fetchProfile()
            if (!profile.error) {
                next()
            } else {
                redirectToLogin()
            }
        } catch {
            redirectToLogin()
        }
    }

    const requireGuest: Callback = async (_, next) => {
        const redirectToHome = () => router.redirect("/")
        try {
            const profile = await fetchProfile()
            if (profile.error) {
                next()
            } else {
                redirectToHome()
            }
        } catch {
            next()
        }
    }

    const setPage: (c: Component) => Callback = (nextPage) => {
        return (_ctx, _next) => {
            page = nextPage
        }
    }

    const registerRoutes = () => {
        router("/", setPage(Index))
        router("/about", setPage(About))
        router("/login", requireGuest, setPage(Login))
        router("/register", requireGuest, setPage(Register))
        router("/recipes/create", requireLogin, setPage(CreateRecipe))
        router("*", setPage(NotFound))

        router.start()
    }

    return {
        get page() {
            return page
        },
        registerRoutes,
    }
}