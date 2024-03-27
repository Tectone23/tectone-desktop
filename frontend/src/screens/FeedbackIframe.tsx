import Heading from "../components/Heading";
import Container from "./Container";

export default function FeedbackIframe() {
  return (
    <Container>
      <Heading
        title="Feedback"
        description="Share your feedback and ideas about using Tectone Desktop. We highly value and encourage feedback from the Tectone community."
      />
      <div className="w-full flex flex-col max-w-3xl space-y-6 ">
        <script src="https://desk.zoho.eu/portal/api/feedbackwidget/164109000000269001?orgId=20096507766&displayType=iframe"></script>
        <iframe
          className=""
          id="zsfeedbackFrame"
          width="890"
          height="570"
          name="zsfeedbackFrame"
          allowTransparency={false}
          src="https://desk.zoho.eu/support/fbw?formType=AdvancedWebForm&fbwId=edbsn4fe291fb3eb5a12574829ceedb1ff928cd4cb1ad897805ab16559107835d76b8&xnQsjsdp=edbsn9f96b7fcfd5e893af845313cf0b6b444&mode=showNewWidget&displayType=iframe"
        ></iframe>
      </div>
    </Container>
  );
}
