import { useParams } from "react-router-dom";
import { Button, Card, Col, Form, Image, Row } from "react-bootstrap";
import { KeyBoxedContent, KeyPairsSection } from "./styles";
import { PageHeader } from "../../../components/layouts/styles";

// Images
import dsManageImg from "../../../assets/images/ds-manage.png";

const GetResultsPage = () => {
  const params = useParams();

  return (
    <>
      <PageHeader>
        <h2>
          <Image className="titleIcon" src={dsManageImg} />
          <span>Get Results</span>
        </h2>
      </PageHeader>
      <Card className="mt-4">
        <Card.Body>
          <KeyBoxedContent>
            <Row>
              <Col md={12} lg={6}>
                <KeyPairsSection>
                  <h3>Get Results</h3>
                </KeyPairsSection>
              </Col>
            </Row>
            <Row className="mt-4">
              <Col md={8}>
                <Form.Group className="mb-3">
                  <Form.Label>Job ID</Form.Label>
                  <Form.Control
                    type="text"
                    placeholder="Job ID"
                    value={params.id}
                    onChange={() => {
                      console.log("job id");
                    }}
                  />
                </Form.Group>
              </Col>
              <Col md={4}></Col>
            </Row>
            <Row className="mt-4">
              <Col md={8}>
                <Form.Group className="mb-3">
                  <Form.Label>Flags</Form.Label>
                  <Form.Control type="text" placeholder="Flags" />
                </Form.Group>
              </Col>
              <Col md={4}></Col>
            </Row>
            <Row>
              <Col md={8} className="text-left">
                <Button>Get Compute Result</Button>
              </Col>
              <Col md={4}></Col>
            </Row>
          </KeyBoxedContent>
        </Card.Body>
      </Card>
    </>
  );
};

export default GetResultsPage;
