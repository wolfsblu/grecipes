<script lang="ts">
    import Button from "../lib/components/forms/Button.svelte";
    import Input from "../lib/components/forms/Input.svelte";
    import Layout from "../Layout.svelte";
    import LoginIcon from "../lib/icons/Login.svelte";
    import Navbar from "../lib/components/navigation/Navbar.svelte";
    import t from "../lib/i18n/i18n.svelte"
    import {createRouter} from "../lib/router.svelte";
    import {createUser} from "../lib/auth/user.svelte";
    import food from "../assets/images/login.jpg"

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
    <div class="border border-gray-200 grid grid-flow-col grid-cols-[1fr_2fr] h-1/2 rounded-md">
        <div class="bg-center bg-cover rounded-l-md" style="background-image: url({food})">
        </div>
        <div class="p-6">
            <h1 class="font-light mb-3 text-3xl">{t("login.title")}</h1>
            {#if error}
                <p>{error.message}</p>
            {/if}
            <form onsubmit="{handleSubmit}">
                <Input label={t("login.labels.email")} type="email" bind:value={credentials.email} required={true} />
                <Input label={t("login.labels.password")} type="password" bind:value={credentials.password} required={true} />
                <Button type="submit" class="mt-3" icon={LoginIcon}>
                    {t("login.actions.submit")}
                </Button>
            </form>
        </div>
    </div>
</Layout>
