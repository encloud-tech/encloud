import * as Yup from "yup";
import * as formik from "formik";
import { Link, useLocation } from "react-router-dom";
import { Button, Card, Col, Image, Row, Modal, Form } from "react-bootstrap";
import { ColoredBtn, SectionBox } from "./styles";
import { PageHeader } from "../../../components/layouts/styles";

// Images
import dsRefreshImg from "../../../assets/images/ds-refresh.png";
import { CSSProperties, useState } from "react";
import { RetrieveByUUID } from "../../../../wailsjs/go/main/App";
import { readKey } from "../../../services/localStorage.service";
import { toast } from "react-toastify";
import ClipLoader from "react-spinners/ClipLoader";

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

const override: CSSProperties = {
  display: "block",
  margin: "0 auto",
  borderColor: "white",
};

const RetrieveContentPage = () => {
  const [downloadLoading, setDownloadLoading] = useState(false);
  const [open, setOpen] = useState(false);
  const location = useLocation();
  const { metadata } = location.state;

  const { Formik } = formik;

  const schema = Yup.object().shape({
    filePath: Yup.string().required("Please enter download file path"),
  });

  const download = (downloadFilePath: string) => {
    try {
      RetrieveByUUID(
        metadata.uuid,
        readKey().PublicKey,
        readKey().PrivateKey,
        downloadFilePath
      )
        .then((result: any) => {
          if (result && result.Status == "success") {
            setDownloadLoading(false);
            toast.success("Document downloaded successfully.", {
              position: toast.POSITION.TOP_RIGHT,
            });
          }
        })
        .catch((err: any) => {
          setDownloadLoading(false);
          toast.error("Something went wrong!.Please retry", {
            position: toast.POSITION.TOP_RIGHT,
          });
        });
    } catch (err) {
      setDownloadLoading(false);
      toast.error("Something went wrong!.Please retry", {
        position: toast.POSITION.TOP_RIGHT,
      });
    }
  };

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
              <ColoredBtn
                className={`step-button ml-2 ${
                  downloadLoading ? "loadingStatus" : ""
                }`}
                style={{ float: "right" }}
                disabled={downloadLoading}
                onClick={() => {
                  setOpen(true);
                }}
              >
                {downloadLoading ? (
                  <div>
                    <ClipLoader
                      color="#ffffff"
                      loading={downloadLoading}
                      cssOverride={override}
                      size={30}
                      aria-label="Loading Spinner"
                      data-testid="loader"
                    />
                    <span className="loadingText">Downloading</span>
                  </div>
                ) : (
                  "Download"
                )}
              </ColoredBtn>
              <Link
                to="/list"
                className="btn btn-primary step-button"
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
      <Modal show={open} onHide={() => setOpen(false)}>
        <Modal.Header closeButton>
          <Modal.Title>Download Content</Modal.Title>
        </Modal.Header>
        <Formik
          validationSchema={schema}
          onSubmit={(data: any) => {
            download(data.filePath);
            setOpen(false);
          }}
          initialValues={{
            filePath: "",
          }}
        >
          {({
            handleSubmit,
            handleChange,
            handleBlur,
            values,
            touched,
            isValid,
            errors,
          }) => (
            <Form noValidate onSubmit={handleSubmit}>
              <Modal.Body>
                <Form.Control
                  type="text"
                  placeholder="Download File Path"
                  aria-label="Download File Path"
                  name="filePath"
                  value={values.filePath}
                  onChange={handleChange}
                  isInvalid={!!errors.filePath}
                />
                <span
                  className="invalid-feedback"
                  style={{ color: "red", textAlign: "left" }}
                >
                  {errors.filePath}
                </span>
              </Modal.Body>
              <Modal.Footer>
                <Button variant="secondary" onClick={() => setOpen(false)}>
                  Close
                </Button>
                <Button type="submit" variant="primary">
                  Download
                </Button>
              </Modal.Footer>
            </Form>
          )}
        </Formik>
      </Modal>
    </>
  );
};

export default RetrieveContentPage;
