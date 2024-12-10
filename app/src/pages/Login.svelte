<script lang="ts">
    import Layout from "../Layout.svelte";
    import Navbar from "../lib/components/navigation/Navbar.svelte";
    import {createRouter} from "../lib/router.svelte";
    import {createUser} from "../lib/auth/user.svelte";
    import Input from "../lib/components/forms/Input.svelte";
    import Button from "../lib/components/forms/Button.svelte";
    import t from "../lib/i18n/i18n.svelte"

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
            await user.login(credentials)
            router.redirectToNext()
        } catch (e) {
            error = e as Error
        }
    }
</script>

<Layout Header={Navbar}>
    <h1>{t("login.title")}</h1>
    {#if error}
        <p>{error.message}</p>
    {/if}
    <form onsubmit="{handleSubmit}">
        <Input label={t("login.labels.email")} type="email" bind:value={credentials.email} />
        <Input label={t("login.labels.password")} type="password" bind:value={credentials.password} />
        <Button type="submit" class="mt-3">
            {t("login.actions.submit")}
        </Button>
    </form>
</Layout>
