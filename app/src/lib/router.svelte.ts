import router, {type Callback} from "page"
import {type Component} from "svelte"
import About from "../pages/About.svelte"
import CreateRecipe from "../pages/recipes/Create.svelte"
import Index from "../pages/Index.svelte"
import Login from "../pages/Login.svelte";
import NotFound from "../pages/errors/404.svelte"
import Register from "../pages/Register.svelte";
import {createUser} from "./auth/user.svelte";


const user = createUser()

let page: Component | null = $state(null);


export const createRouter = () => {
    const nextParam = "next"

    const redirect = (to: string) => router.redirect(to)
    const redirectToHome = () => redirect("/")
    const redirectToLogin = (nextRoute: string) => redirect(`/login?${nextParam}=${encodeURIComponent(nextRoute)}`)
    const redirectToNext = () => {
        const queryParams = new URLSearchParams(window.location.search)
        const next = queryParams.get(nextParam)
        if (next) {
            router.redirect(next)
        }
    }

    const requireLogin: Callback = async (ctx, next) => {
        if (user.profile) {
            next()
        } else {
            redirectToLogin(ctx.path)
        }
    }

    const requireGuest: Callback = async (_, next) => {
        if (user.profile) {
            redirectToHome()
        } else {
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
        redirectToNext,
        registerRoutes,
    }
}