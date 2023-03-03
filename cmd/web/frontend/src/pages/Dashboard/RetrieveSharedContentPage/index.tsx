import * as Yup from "yup";
import * as formik from "formik";
import { Button, Card, Col, Form, Image, Row } from "react-bootstrap";
import { KeyBoxedContent, KeyPairsSection } from "./styles";
import { PageHeader } from "../../../components/layouts/styles";

// Images
import dsManageImg from "../../../assets/images/ds-manage.png";
import { RetrieveSharedContent } from "../../../../wailsjs/go/main/App";
import { toast } from "react-toastify";
import { CSSProperties, useState } from "react";
import ClipLoader from "react-spinners/ClipLoader";
import { ColoredBtn } from "./styles";

const override: CSSProperties = {
  display: "block",
  margin: "0 auto",
  borderColor: "white",
};

const RetrieveSharedContentPage = () => {
  const [loading, setLoading] = useState(false);
  const { Formik } = formik;

  const schema = Yup.object().shape({
    cid: Yup.string().required("Please enter cid"),
    dekType: Yup.string().required("Please enter dek type"),
    decryptedDekPath: Yup.string().required("Please enter dek path"),
    fileName: Yup.string().required("Please enter download file name"),
    retrievalFileStoragePath: Yup.string().required(
      "Please enter download file path"
    ),
  });

  const getSharedContent = (data: any) => {
    setLoading(true);
    try {
      RetrieveSharedContent(
        data.decryptedDekPath,
        data.dekType,
        data.cid,
        data.fileName,
        data.retrievalFileStoragePath
      )
        .then((result: any) => {
          if (result && result.Status == "success") {
            setLoading(false);
            toast.success("Document downloaded successfully.", {
              position: toast.POSITION.TOP_RIGHT,
            });
          }
        })
        .catch((err: any) => {
          setLoading(false);
          toast.error("Something went wrong!.Please retry", {
            position: toast.POSITION.TOP_RIGHT,
          });
        });
    } catch (err) {
      setLoading(false);
      toast.error("Something went wrong!.Please retry", {
        position: toast.POSITION.TOP_RIGHT,
      });
    }
  };

  return (
    <>
      <PageHeader>
        <h2>
          <Image className="titleIcon" src={dsManageImg} />
          <span>Retrieve Shared Content</span>
        </h2>
      </PageHeader>
      <Card className="mt-4">
        <Card.Body>
          <KeyBoxedContent>
            <Formik
              validationSchema={schema}
              onSubmit={getSharedContent}
              initialValues={{
                cid: "",
                dekType: "",
                decryptedDekPath: "",
                fileName: "",
                retrievalFileStoragePath: "",
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
                  <Row>
                    <Col md={12} lg={6}>
                      <KeyPairsSection>
                        <h3>Retrieve Shared Content</h3>
                      </KeyPairsSection>
                    </Col>
                  </Row>
                  <Row className="mt-4">
                    <Col md={8}>
                      <Form.Group className="mb-3">
                        <Form.Label>CID</Form.Label>
                        <Form.Control
                          type="text"
                          placeholder="CID"
                          name="cid"
                          value={values.cid}
                          onChange={handleChange}
                          isInvalid={!!errors.cid}
                        />
                        <span
                          className="invalid-feedback"
                          style={{ color: "red", textAlign: "left" }}
                        >
                          {errors.cid}
                        </span>
                      </Form.Group>
                    </Col>
                    <Col md={4}></Col>
                  </Row>
                  <Row className="mt-4">
                    <Col md={8}>
                      <Form.Group className="mb-3">
                        <Form.Label>DEK Type</Form.Label>
                        <Form.Select
                          name="dekType"
                          value={values.dekType}
                          onChange={handleChange}
                          isInvalid={!!errors.dekType}
                        >
                          <option>Please select dek type</option>
                          <option value="aes">AES 256 GCM</option>
                          <option value="chacha20">ChaCha20-Poly1305</option>
                        </Form.Select>
                        <span
                          className="invalid-feedback"
                          style={{ color: "red", textAlign: "left" }}
                        >
                          {errors.dekType}
                        </span>
                      </Form.Group>
                    </Col>
                    <Col md={4}></Col>
                  </Row>
                  <Row className="mt-4">
                    <Col md={8}>
                      <Form.Group className="mb-3">
                        <Form.Label>DEK Path</Form.Label>
                        <Form.Control
                          type="text"
                          placeholder="DEK Path"
                          name="decryptedDekPath"
                          value={values.decryptedDekPath}
                          onChange={handleChange}
                          isInvalid={!!errors.decryptedDekPath}
                        />
                        <span
                          className="invalid-feedback"
                          style={{ color: "red", textAlign: "left" }}
                        >
                          {errors.decryptedDekPath}
                        </span>
                      </Form.Group>
                    </Col>
                    <Col md={4}></Col>
                  </Row>
                  <Row className="mt-4">
                    <Col md={8}>
                      <Form.Group className="mb-3">
                        <Form.Label>Download File Name</Form.Label>
                        <Form.Control
                          type="text"
                          placeholder="sample.csv"
                          name="fileName"
                          value={values.fileName}
                          onChange={handleChange}
                          isInvalid={!!errors.fileName}
                        />
                        <span
                          className="invalid-feedback"
                          style={{ color: "red", textAlign: "left" }}
                        >
                          {errors.fileName}
                        </span>
                      </Form.Group>
                    </Col>
                    <Col md={4}></Col>
                  </Row>
                  <Row className="mt-4">
                    <Col md={8}>
                      <Form.Group className="mb-3">
                        <Form.Label>Download File Path</Form.Label>
                        <Form.Control
                          type="text"
                          placeholder="C:\Users\{YOUR_USERNAME}\Downloads"
                          name="retrievalFileStoragePath"
                          value={values.retrievalFileStoragePath}
                          onChange={handleChange}
                          isInvalid={!!errors.retrievalFileStoragePath}
                        />
                        <span
                          className="invalid-feedback"
                          style={{ color: "red", textAlign: "left" }}
                        >
                          {errors.retrievalFileStoragePath}
                        </span>
                      </Form.Group>
                    </Col>
                    <Col md={4}></Col>
                  </Row>
                  <Row>
                    <Col md={8} className="text-left">
                      <ColoredBtn
                        className={`step-button ml-2 ${
                          loading ? "loadingStatus" : ""
                        }`}
                        disabled={loading}
                        onClick={handleSubmit}
                      >
                        {loading ? (
                          <div>
                            <ClipLoader
                              color="#ffffff"
                              loading={loading}
                              cssOverride={override}
                              size={30}
                              aria-label="Loading Spinner"
                              data-testid="loader"
                            />
                            <span className="loadingText">Downloading</span>
                          </div>
                        ) : (
                          "Get Shared Content"
                        )}
                      </ColoredBtn>
                    </Col>
                    <Col md={4}></Col>
                  </Row>
                </Form>
              )}
            </Formik>
          </KeyBoxedContent>
        </Card.Body>
      </Card>
    </>
  );
};

export default RetrieveSharedContentPage;
