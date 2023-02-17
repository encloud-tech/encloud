import { Link, useLocation } from "react-router-dom";
import { Button, Card, Col, Form, Image, Row } from "react-bootstrap";
import { SectionBox } from "./styles";
import { PageHeader } from "../../../components/layouts/styles";

// Images
import dsRefreshImg from "../../../assets/images/ds-refresh.png";

const dekTypesNames = [
  {
    value: "aes",
    name: "AES 256 GCM",
  },
  {
    value: "chacha20",
    name: "ChaCha20-Poly1305",
  },
];

const RetrieveContentPage = () => {
  const location = useLocation();
  const { metadata } = location.state;

  return (
    <>
      <PageHeader>
        <h2>
          <Image className="titleIcon" src={dsRefreshImg} />
          <span>Retrieve Content</span>
        </h2>
      </PageHeader>
      <SectionBox>
        <Card className="mt-4">
          <Card.Header>
            <h4>
              Content Details{" "}
              <Button style={{ float: "right" }}>Download Content</Button>
              <Link
                to="/list"
                className="btn btn-primary"
                style={{ marginRight: 5, float: "right" }}
              >
                Back
              </Link>
            </h4>
          </Card.Header>
          <Card.Body>
            <Row>
              <Col sm="3">
                <Card.Text className="fw-bold">UUID</Card.Text>
              </Col>
              <Col sm="9">
                <Card.Text className="text-muted">{metadata.uuid}</Card.Text>
              </Col>
            </Row>
            <hr />
            <Row>
              <Col sm="3">
                <Card.Text className="fw-bold">File Name</Card.Text>
              </Col>
              <Col sm="9">
                <Card.Text className="text-muted">{metadata.name}</Card.Text>
              </Col>
            </Row>
            <hr />
            <Row>
              <Col sm="3">
                <Card.Text className="fw-bold">File Size</Card.Text>
              </Col>
              <Col sm="9">
                <Card.Text className="text-muted">{metadata.size}</Card.Text>
              </Col>
            </Row>
            <hr />
            <Row>
              <Col sm="3">
                <Card.Text className="fw-bold">File Type</Card.Text>
              </Col>
              <Col sm="9">
                <Card.Text className="text-muted">
                  {metadata.fileType}
                </Card.Text>
              </Col>
            </Row>
            <hr />
            <Row>
              <Col sm="3">
                <Card.Text className="fw-bold">CID</Card.Text>
              </Col>
              <Col sm="9">
                <Card.Text className="text-muted">{metadata.cid}</Card.Text>
              </Col>
            </Row>
            <hr />
            <Row>
              <Col sm="3">
                <Card.Text className="fw-bold">Uploaded At</Card.Text>
              </Col>
              <Col sm="9">
                <Card.Text className="text-muted">
                  {metadata.uploadedAt}
                </Card.Text>
              </Col>
            </Row>
            <hr />
            <Row>
              <Col sm="3">
                <Card.Text className="fw-bold">Dek Type</Card.Text>
              </Col>
              <Col sm="9">
                <Card.Text className="text-muted">
                  {
                    dekTypesNames.find((o) => o.value === metadata.dekType)
                      ?.name
                  }
                </Card.Text>
              </Col>
            </Row>
            <hr />
            <Row>
              <Col sm="3">
                <Card.Text className="fw-bold">Kek Type</Card.Text>
              </Col>
              <Col sm="9">
                <Card.Text className="text-muted">
                  {metadata.kekType.toUpperCase()}
                </Card.Text>
              </Col>
            </Row>
          </Card.Body>
        </Card>
      </SectionBox>
    </>
  );
};

export default RetrieveContentPage;
