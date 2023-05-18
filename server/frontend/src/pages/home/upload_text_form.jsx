import React, { useContext, useState } from "react";
import {
  BigTextarea,
  Button,
  Form,
  showUploadingDialog,
  showUploadTextSuccessDialog,
  uploadText,
} from "./components";
import { AppContext } from "../../shared/app_context";
import { Center } from "../../components/center";

export const UploadTextForm = () => {
  const context = useContext(AppContext);
  const [text, setText] = useState("");
  const onSubmit = async (e) => {
    e.preventDefault();
    showUploadingDialog();
    const { data: { url } } = await uploadText(text)
    const port = import.meta.env.VITE_BACKEND_PORT;
    showUploadTextSuccessDialog({
      context, content: (addr) => {
        return addr && `http://${addr}:${port}/static/downloads?type=text&url=http://${addr + `:${port}/api/v1` + encodeURIComponent(url)}`
      }
    });
  };
  return (
    <Form className="uploadForm" onSubmit={onSubmit}>
      <div className="row">
        <BigTextarea
          value={text}
          onChange={(e) => setText(e.target.value)}
        />
      </div>
      <Center className="row">
        <Button type="submit">上传</Button>
      </Center>
    </Form>
  );
};
