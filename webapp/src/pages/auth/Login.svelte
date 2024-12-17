<script lang="ts">
    import Input from "../../lib/components/forms/Input.svelte";
    import Layout from "../../Layout.svelte";
    import LoginIcon from "../../lib/icons/Login.svelte";
    import Navbar from "../../lib/components/navigation/Navbar.svelte";
    import t from "../../lib/i18n/i18n.svelte.js"
    import {createRouter} from "../../lib/router.svelte.js";
    import {createUser} from "../../lib/auth/user.svelte.js";
    import food from "../../assets/images/auth/login.jpg"
    import Form from "../../lib/components/auth/Form.svelte";

    const router = createRouter()
    const user = createUser()

    let credentials: Credentials = $state({
        email: "",
        password: "",
    })

    const handleSubmit = async (e: SubmitEvent) => {
        e.preventDefault()
        try {
            await user.login(credentials)
            router.redirectToNext()
        } catch (e) {
            // TODO: Show error toast
        }
    }
</script>

<Layout Header={Navbar}>
    <div class="flex flex-col flex-wrap md:h-full items-center justify-center">
        <Form imageSrc={food}
              submitIcon={LoginIcon}
              submitLabel={t("login.actions.submit")}
              onSubmit={handleSubmit}
              title={t("login.title")}
              subtitle={t("login.subtitle", {url: '/register'})}
        >
            <Input inputClass="mb-3" label={t("login.labels.email")} type="email" bind:value={credentials.email}
                   required={true}/>
            <Input label={t("login.labels.password")} type="password"
                   bind:value={credentials.password} required={true}/>
            <p class="text-sm">
                {@html t("login.help.resetPassword", {
                    url: "/forgot-password",
                })}
            </p>
        </Form>
    </div>
</Layout>