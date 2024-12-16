import router, {type Callback} from "page"
import {type Component} from "svelte"
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
        const next = queryParams.get(nextParam) || "/"
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

    const setPage: (importFn: () => Promise<Component>) => Callback = (importFn) => {
        return async (_ctx, _next) => {
            page = await importFn()
        };
    }

    const registerRoutes = async () => {
        router("/", setPage(async () => (await import("../pages/Index.svelte")).default))
        router("/about", setPage(async () => (await import("../pages/About.svelte")).default))
        router("/login", requireGuest, setPage(async () => (await import("../pages/auth/Login.svelte")).default))
        router("/recipes/create", requireLogin, setPage(async () => (await import("../pages/recipes/Create.svelte")).default))
        router("/register", requireGuest, setPage(async () => (await import("../pages/auth/Register.svelte")).default))
        router("/forgot-password", requireGuest, setPage(async () => (await import("../pages/auth/ForgotPassword.svelte")).default))
        router("*", setPage(async () => (await import("../pages/errors/404.svelte")).default))

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