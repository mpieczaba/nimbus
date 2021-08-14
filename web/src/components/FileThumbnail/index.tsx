import React from "react";

import { Wrapper, Thumbnail } from "./styles";

import FileIcon from "../FileIcon";

interface Props {
  extension: string;
  url?: string;
  image?: boolean;
}

const FileThumbnail: React.FC<Props> = ({ extension, url, image }) => {
  return (
    <Wrapper>
      {hasThumbnail(extension) && url && image ? (
        <Thumbnail src={url} />
      ) : (
        <FileIcon icon={getIcon(extension)} />
      )}
    </Wrapper>
  );
};

const hasThumbnail = (extension: string): boolean => {
  // https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types#image_types
  switch (extension) {
    case ".apng":
    case ".avif":
    case ".gif":
    case ".jpg":
    case ".jpeg":
    case ".jfif":
    case ".pjpeg":
    case ".pjp":
    case ".png":
    case ".svg":
    case ".webp":
      return true;
    default:
      return false;
  }
};

const getIcon = (extension: string): string => {
  switch (extension) {
    case ".as":
      return "actionscript";
    case ".apk":
      return "android";
    // return "angular"
    // return "arduino"
    case ".s":
    case ".asm":
      return "assembly";
    case ".flac":
    case ".mp3":
    case ".ogg":
    case ".wav":
      return "audio";
    // return "authors"
    case ".c":
      return "c";
    // return "changelog"
    // return "conduct"
    // return "console"
    // return "contributing"
    case ".cpp":
      return "cpp";
    // return "credits"
    case ".cs":
      return "csharp";
    case ".css":
      return "css3";
    case ".d":
      return "d";
    case ".dart":
      return "dart";
    case ".sql":
      return "database";
    case ".txt":
      return "document";
    case ".ex":
    case ".exs":
      return "elixir";
    // return "email"
    case ".erl":
    case ".hrl":
      return "erlang";
    case ".exe":
      return "exe";
    case ".woff":
    case ".ttf":
    case ".otf":
      return "font";
    case ".fs":
      return "fsharp";
    // return "git"
    // return "gitlab"
    case ".go":
      return "go";
    case ".graphql":
      return "graphql";
    case ".h":
      return "h";
    case ".haml":
      return "haml";
    case ".hs":
    case ".lhs":
      return "haskell";
    case ".hpp":
      return "hpp";
    case ".html":
    case ".htm":
      return "html";
    // return "http"
    case ".apng":
    case ".avif":
    case ".gif":
    case ".jpg":
    case ".jpeg":
    case ".jfif":
    case ".pjpeg":
    case ".pjp":
    case ".png":
    case ".svg":
    case ".webp":
      return "image";
    case ".java":
    case ".class":
    case ".jar":
      return "java";
    case ".js":
      return "javascript";
    case ".json":
      return "json";
    case ".jl":
      return "julia";
    case ".key":
      return "key";
    case ".kt":
    case ".kts":
    case ".ktm":
      return "kotlin";
    case ".less":
      return "less";
    // return "lib"
    case ".ls":
      return "livescript";
    case ".lock":
      return "lock";
    case ".log":
      return "log";
    case ".lua":
      return "lua";
    case ".md":
      return "markdown";
    case ".adp":
    case ".adn":
      return "microsoft-access";
    case ".xlsx":
    case ".xlsm":
      return "microsoft-excel";
    case ".onetoc":
    case ".one":
      return "microsoft-onenote";
    case ".pptx":
    case ".ppt":
      return "microsoft-powerpoint";
    case ".docx":
    case ".doc":
      return "microsoft-word";
    case ".mxml":
      return "mxml";
    case ".ml":
    case ".mli":
      return "ocaml";
    case ".pdf":
      return "pdf";
    case ".plx":
    case ".pl":
      return "perl";
    case ".php":
      return "php";
    case ".pug":
      return "pug";
    case ".purs":
      return "purescript";
    case ".py":
      return "python";
    case ".r":
    case ".rdata":
      return "r";
    case ".raml":
      return "raml";
    case ".jsx":
    case ".tsx":
      return "react";
    case ".rb":
      return "ruby";
    case ".rs":
      return "rust";
    case ".sass":
    case ".scss":
      return "sass";
    case ".scala":
    case ".sc":
      return "scala";
    // return "settings"
    case ".swift":
      return "swift";
    case ".ts":
      return "typescript";
    // return "url"
    case ".3gp":
    case ".mp4":
    case ".m4v":
    case ".m4p":
    case ".mov":
    case ".webm":
      return "video";
    // return "virtual"
    case ".vue":
      return "vue";
    case ".wasm":
      return "webassembly";
    case ".xaml":
      return "xaml";
    case ".xml":
      return "xml";
    case ".yaml":
      return "yaml";
    case ".zip":
      return "zip";
    default:
      return "file";
  }
};

export default FileThumbnail;
