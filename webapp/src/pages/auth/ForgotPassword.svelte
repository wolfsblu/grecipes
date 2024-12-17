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

    let email = $state("")

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
              submitLabel={t("forgot-password.actions.submit")}
              onSubmit={handleSubmit}
              title={t("forgot-password.title")}
              subtitle={t("forgot-password.subtitle")}
        >
            <Input label={t("forgot-password.labels.email")} type="email" bind:value={email} required={true}/>
            <p class="text-sm">
                {@html t("forgot-password.help.email", {
                    class: "hover:no-underline hover:text-blue-500 text-blue-600 underline",
                    url: "/"
                })}
            </p>
        </Form>
    </div>
</Layout>
