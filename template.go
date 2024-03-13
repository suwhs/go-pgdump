package pgdump

import (
	"os"
	"text/template"
)

type DumpInfo struct {
	DumpVersion   string
	ServerVersion string
	CompleteTime  string
}

func writeHeader(file *os.File, info DumpInfo) error {
    const headerTemplate = `-- Go PostgreSQL Dump {{ .DumpVersion }}
--
-- ------------------------------------------------------
-- Server version    {{ .ServerVersion }}

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;
`
    tmpl, err := template.New("header").Parse(headerTemplate)
    if err != nil {
        return err
    }
    return tmpl.Execute(file, info)
}

func writeFooter(file *os.File, info DumpInfo) error {
    const footerTemplate = `--
-- Dump completed on {{ .CompleteTime }}
--`
    tmpl, err := template.New("footer").Parse(footerTemplate)
    if err != nil {
        return err
    }
    return tmpl.Execute(file, info)
}
