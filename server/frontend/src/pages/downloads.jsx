import ClipboardJS from 'clipboard';
import { useQuery } from "../hooks/use_query";
import { BigTextarea, Button, Header, Layout } from "./home/components";
import styled from "styled-components";
import { Center } from "../components/center";
import { Space } from "../components/space";
import { useState } from "react";
import { useRef } from "react";
import { useEffect } from "react";
import { http } from "../shared/http";
import { history } from "../shared/history";

export const Downloads = () => {
  const query = useQuery()
  const type = normalizeType(query.type)
  const [text, setText] = useState("")
  useEffect(() => {
    (async () => {
      if(type !== "text") return;
      const res = await http.get(query.url)
      setText(res)
    })()
  }, [type])
  const handleCopyClick = () => {
    if(!text) return;
    const clipboard = new ClipboardJS('.copy-button', {
      text: () => text
    });
    clipboard.on('success', () => {
      alert("复制成功")
      clipboard.destroy();
    });
    clipboard.on('error', () => {
      alert("复制失败")
      clipboard.destroy();
    });

    clipboard.onClick({ trigger: document.querySelector('.copy-button') });
  };
  let node = null
  switch (type) {
    case 'text':
      node = (
        <div>
          <BigTextarea readOnly value={text} />
          <Space />
          <Center virtical>
            <Button className="copy-button" onClick={handleCopyClick}>点击复制</Button>
          </Center>
        </div>
      )
      break;
    case 'file':
      node = (
        <Center virtical>
          <a href={query.url}>
            <svg><use xlinkHref="#icon-file"></use></svg>
            <Space />
            <Center>
              <Button>点击下载文件</Button>
            </Center>
          </a>
        </Center>
      )
      break;
    case 'image':
      node = (
        <Center>
          <a href={query.url}>
            <Picture src={query.url} />
            <Center>
              <Button>长按或点击，即可下载图片</Button>
            </Center>
          </a>
        </Center>
      )
      break;
  }
  const onClickUpload = () => {
    history.push("/")
  }
  return (
    <Layout>
      <Header>同步传</Header>
      {node}
      <Space x3 />
      <Center>
        <Button onClick={onClickUpload}>我也要上传</Button>
      </Center>
    </Layout>
  )
};
const Picture = styled.img`
  border: 2px solid ${({ theme }) => theme.borderColor};
  margin: 16px;
`
const P = styled.p`
  margin: 8px 0;
`

const normalizeType = type => {
  if (/^image\/.*$/.test(type)) {
    return 'image'
  } else if (/^text(\/.*)?$/.test(type)) {
    return 'text'
  } else {
    return 'file'
  }
}