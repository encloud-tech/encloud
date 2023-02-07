import { Button, Card, Col, Form, Image, Row } from "react-bootstrap";
import { KeyBoxedContent, KeyPairsSection } from "./styles";
import { PageHeader } from "./../../../components/layouts/styles";

// Images
import dsManageImg from "../../../assets/images/ds-manage.png";
import { useState } from "react";

const ManageComputePage = () => {
  const [computeSection, setComputeSection] = useState(true);
  const [, setGetResults] = useState(false);

  return (
    <>
      <PageHeader>
        <h2>
          <Image className="titleIcon" src={dsManageImg} />
          <span>Manage Compute</span>
        </h2>
      </PageHeader>
      <Row xs={1} md={2} className="g-4">
        <Col md={3}></Col>
        <Col md={3}>
          <KeyBoxedContent className={computeSection ? "active" : ""}>
            <Card>
              <Card.Body>
                <KeyPairsSection>
                  <h3
                    style={{ cursor: "pointer" }}
                    onClick={() => {
                      setComputeSection(true);
                      setGetResults(false);
                    }}
                  >
                    Config Compute
                  </h3>
                </KeyPairsSection>
              </Card.Body>
            </Card>
          </KeyBoxedContent>
        </Col>
        <Col md={3}>
          <KeyBoxedContent className={!computeSection ? "active" : ""}>
            <Card>
              <Card.Body>
                <KeyPairsSection>
                  <h3
                    style={{ cursor: "pointer" }}
                    onClick={() => {
                      setGetResults(true);
                      setComputeSection(false);
                    }}
                  >
                    Get Results
                  </h3>
                </KeyPairsSection>
              </Card.Body>
            </Card>
          </KeyBoxedContent>
        </Col>
        <Col md={3}></Col>
      </Row>

      <Card className="mt-4">
        <Card.Body>
          <KeyBoxedContent>
            <Row>
              <Col md={12} lg={6}>
                <KeyPairsSection>
                  <h3>{computeSection ? "Config Compute" : "Get Results"}</h3>
                </KeyPairsSection>
              </Col>
            </Row>
            {computeSection ? (
              <>
                <Row className="mt-4">
                  <Col md={8}>
                    <Form.Group className="mb-3">
                      <Form.Label>Flags</Form.Label>
                      <Form.Control type="text" placeholder="Flags" />
                    </Form.Group>
                  </Col>
                  <Col md={4}></Col>
                </Row>
                <Row className="mt-4">
                  <Col md={8}>
                    <Form.Group className="mb-3">
                      <Form.Label>Docker Image URI</Form.Label>
                      <Form.Control
                        type="text"
                        placeholder="Docker Image URI"
                      />
                    </Form.Group>
                  </Col>
                  <Col md={4}></Col>
                </Row>
                <Row className="mt-4">
                  <Col md={8}>
                    <Form.Group className="mb-3">
                      <Form.Label>Commands</Form.Label>
                      <Form.Control
                        type="text"
                        placeholder="Commands and args"
                      />
                    </Form.Group>
                  </Col>
                  <Col md={4}></Col>
                </Row>
              </>
            ) : (
              <>
                <Row className="mt-4">
                  <Col md={8}>
                    <Form.Group className="mb-3">
                      <Form.Label>Job ID</Form.Label>
                      <Form.Control type="text" placeholder="Job ID" />
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
              </>
            )}
            <Row>
              <Col md={8} className="text-left">
                <Button>
                  {computeSection ? "Run Compute" : "Get Compute Results"}
                </Button>
              </Col>
              <Col md={4}></Col>
            </Row>
          </KeyBoxedContent>
        </Card.Body>
      </Card>
    </>
  );
};

export default ManageComputePage;
