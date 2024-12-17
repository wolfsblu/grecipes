<script lang="ts">
    import Input from "../../lib/components/forms/Input.svelte";
    import Layout from "../../Layout.svelte";
    import SubmitIcon from "../../lib/icons/Submit.svelte";
    import Navbar from "../../lib/components/navigation/Navbar.svelte";
    import t from "../../lib/i18n/i18n.svelte.js"
    import {createRouter} from "../../lib/router.svelte.js";
    import {createUser} from "../../lib/auth/user.svelte.js";
    import food from "../../assets/images/auth/forgot.jpg"
    import Form from "../../lib/components/auth/Form.svelte";

    const router = createRouter()
    const user = createUser()

    let password = $state("")

    const handleSubmit = async (e: SubmitEvent) => {
        e.preventDefault()
        try {
            await user.resetPassword(email)
            router.redirectToNext()
        } catch (e) {
            // TODO: Show error toast
        }
    }
</script>

<Layout Header={Navbar}>
    <div class="flex flex-col flex-wrap md:h-full items-center justify-center">
        <Form imageSrc={food}
              submitIcon={SubmitIcon}
              submitLabel={t("reset-password.actions.submit")}
              onSubmit={handleSubmit}
              title={t("reset-password.title")}
              subtitle={t("reset-password.subtitle")}
        >
            <div class="gap-3 grid grid-cols-1 md:grid-cols-2">
                <Input label={t("reset-password.labels.password")} type="password"
                       bind:value={password} required={true}/>
                <Input label={t("reset-password.labels.confirmPassword")} type="password"
                       bind:value={password} required={true}/>
            </div>
        </Form>
    </div>
</Layout>
