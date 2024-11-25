<script lang="ts">
    import Layout from "../Layout.svelte";
    import Navbar from "../lib/components/navigation/Navbar.svelte";
    import {login} from "../lib/api/client";
    import {createRouter} from "../lib/router.svelte";
    import {createUser} from "../lib/auth/user.svelte";

    const router = createRouter()
    const user = createUser()

    let credentials: Credentials = $state({
        email: "",
        password: "",
    })

    let error: Error | null = $state(null)

    const handleSubmit = async (e: SubmitEvent) => {
        e.preventDefault()
        error = null
        try {
            const response = await login(credentials)
            if (response.error) {
                error = response.error as Error
            } else {
                user.login(response.data)
            }
        } catch (e) {
            error = e as Error
        }
        const queryParams = new URLSearchParams(window.location.search)
        const next = queryParams.get("next")
        if (next) {
            router.redirect(next)
        }
    }
</script>

<Layout Header={Navbar}>
    <h1>Login</h1>
    {#if error}
        <p>{error.message}</p>
    {/if}
    <form onsubmit="{handleSubmit}">
        <input bind:value={credentials.email} type="email">
        <input bind:value={credentials.password} type="password">

        <button type="submit">
            Login
        </button>
    </form>
</Layout>
