import { Button, Card, Col, Form, Image, Row } from "react-bootstrap";
import { KeyBoxedContent, KeyPairsSection } from "./styles";
import { PageHeader } from "./../../../components/layouts/styles";

// Images
import dsManageImg from "../../../assets/images/ds-manage.png";
import { useEffect, useState } from "react";

const ManageComputePage = () => {
  const [preview, setPreview] = useState("");
  const [values, setValues] = useState({
    flags: "",
    dockerImageURI: "",
    commands: "",
  });

  const handleChange = (event: any) => {
    setValues((values) => ({
      ...values,
      [event.target.name]: event.target.value,
    }));
  };

  useEffect(() => {
    if (values && (values.flags || values.dockerImageURI || values.commands)) {
      setPreview(
        "bacalhau docker run " +
          values.flags +
          " " +
          values.dockerImageURI +
          " " +
          values.commands
      );
    } else {
      setPreview("");
    }
  }, [values]);

  return (
    <>
      <PageHeader>
        <h2>
          <Image className="titleIcon" src={dsManageImg} />
          <span>Manage Compute</span>
        </h2>
      </PageHeader>
      <Card className="mt-4">
        <Card.Body>
          <KeyBoxedContent>
            <Row>
              <Col md={12} lg={6}>
                <KeyPairsSection>
                  <h3>Config Compute</h3>
                </KeyPairsSection>
              </Col>
            </Row>
            <Row className="mt-4">
              <Col md={8}>
                <Form.Group className="mb-3">
                  <Form.Label>Flags</Form.Label>
                  <Form.Control
                    type="text"
                    placeholder="-v QmeZRGhe4PmjctYVSVHuEiA9oSXnqmYa4kQubSHgWbjv72:/input_images"
                    name="flags"
                    value={values.flags}
                    onChange={handleChange}
                  />
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
                    placeholder="dpokidov/imagemagick:7.1.0-47-ubuntu"
                    name="dockerImageURI"
                    value={values.dockerImageURI}
                    onChange={handleChange}
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
                    placeholder="magick mogrify -resize 100x100 -quality 100"
                    name="commands"
                    value={values.commands}
                    onChange={handleChange}
                  />
                </Form.Group>
              </Col>
              <Col md={4}></Col>
            </Row>
            <Row className="mt-4">
              <Col md={8}>
                <Form.Group className="mb-3">
                  <Form.Label>Preview</Form.Label>
                  <Form.Control
                    type="text"
                    placeholder="bacalhau docker run [flags] IMAGE[:TAG|@DIGEST] [COMMAND] [ARG...]"
                    value={preview}
                  />
                </Form.Group>
              </Col>
              <Col md={4}></Col>
            </Row>
            <Row>
              <Col md={8} className="text-left">
                <Button>Run Compute</Button>
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
