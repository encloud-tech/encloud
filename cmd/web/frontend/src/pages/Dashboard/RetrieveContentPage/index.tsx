import { Link, useLocation } from "react-router-dom";
import { Button, Card, Col, Image, Row, Modal, Form } from "react-bootstrap";
import { ColoredBtn, SectionBox } from "./styles";
import { PageHeader } from "../../../components/layouts/styles";

// Images
import dsRefreshImg from "../../../assets/images/refresh.png";
import { CSSProperties, useState } from "react";
import {
  RetrieveByUUID,
  SelectDirectory,
} from "../../../../wailsjs/go/main/App";
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
  const [showDownloadForm, setShowDownloadForm] = useState(false);
  const location = useLocation();
  const { metadata } = location.state;

  const [selectedDirectory, setSelectedDirectory] = useState<string>();

  const getPath = async (evt: any) => {
    evt.preventDefault();
    try {
      SelectDirectory()
        .then((result: any) => {
          var dt = new DataTransfer();
          dt.items.add(new File([], result));
          evt.target.files = dt.files;
          setSelectedDirectory(result);
        })
        .catch((err: any) => {
          console.error(err);
        });
    } catch (err) {
      console.error(err);
    }
  };

  const download = () => {
    setDownloadLoading(true);
    try {
      if (selectedDirectory) {
        RetrieveByUUID(
          metadata.uuid,
          readKey().PublicKey,
          readKey().PrivateKey,
          selectedDirectory
        )
          .then((result: any) => {
            if (result && result.Status == "success" && result.Data.uuid) {
              setDownloadLoading(false);
              setShowDownloadForm(!showDownloadForm);
              toast.success("Document downloaded successfully.", {
                position: toast.POSITION.TOP_RIGHT,
              });
            } else {
              setDownloadLoading(false);
              toast.error("Something went wrong!.Please retry", {
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
      }
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
              {!showDownloadForm && (
                <Button
                  style={{ float: "right" }}
                  onClick={() => setShowDownloadForm(!showDownloadForm)}
                >
                  Download
                </Button>
              )}
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
            {showDownloadForm ? (
              <>
                <Row className="mb-3">
                  <Col>
                    <span className="fw-bold">
                      {downloadLoading ? "Download is in progress..." : ""}
                    </span>
                  </Col>
                </Row>
                <Row>
                  <Col md={8}>
                    <Form.Control
                      type="file"
                      placeholder="Download File Path"
                      aria-label="Download File Path"
                      name="directoryPath"
                      onClick={getPath}
                    />
                  </Col>
                  <Col md={4}></Col>
                </Row>
                <Row className="mt-3">
                  <Col md={8}>
                    <ColoredBtn
                      className="step-button"
                      onClick={() => setShowDownloadForm(!showDownloadForm)}
                      style={{ marginRight: 2 }}
                    >
                      Close
                    </ColoredBtn>
                    <ColoredBtn
                      className={`step-button ml-2 ${
                        downloadLoading ? "loadingStatus" : ""
                      }`}
                      disabled={downloadLoading || !selectedDirectory}
                      onClick={download}
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
                  </Col>
                  <Col md={4}></Col>
                </Row>
              </>
            ) : (
              <>
                <Row>
                  <Col sm="3">
                    <Card.Text className="fw-bold">UUID</Card.Text>
                  </Col>
                  <Col sm="9">
                    <Card.Text className="text-muted">
                      {metadata.uuid}
                    </Card.Text>
                  </Col>
                </Row>
                <hr />
                <Row>
                  <Col sm="3">
                    <Card.Text className="fw-bold">File Name</Card.Text>
                  </Col>
                  <Col sm="9">
                    <Card.Text className="text-muted">
                      {metadata.name}
                    </Card.Text>
                  </Col>
                </Row>
                <hr />
                <Row>
                  <Col sm="3">
                    <Card.Text className="fw-bold">File Size</Card.Text>
                  </Col>
                  <Col sm="9">
                    <Card.Text className="text-muted">
                      {metadata.size}
                    </Card.Text>
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
              </>
            )}
          </Card.Body>
        </Card>
      </SectionBox>
    </>
  );
};

export default RetrieveContentPage;
