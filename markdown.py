import os
import mistune
import argparse
import glob

from pygments import highlight
from pygments.lexers import get_lexer_by_name
from pygments.formatters import HtmlFormatter

class HighlightRenderer(mistune.Renderer):
    def block_code(self, code, lang):
        if not lang:
            return '\n<pre><code>{}</code></pre>\n'.format(mistune.escape(code))

        lexer = get_lexer_by_name("python", stripall=True)
        formatter = HtmlFormatter(linenos=True, cssclass="source")
        html = highlight(code, lexer, formatter)

        return '\n<pre><code class="{}">{}</code></pre>\n'.format(lang, html)


def markdown2html(path):
    renderer = HighlightRenderer()
    markdown_parser = mistune.Markdown(renderer=renderer)

    with open(path) as markdown:
        markdown_data = markdown.read()

    html = markdown_parser(markdown_data)

    return html


def main(args):
    html = markdown2html(args.markdown)

    print(html)

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('markdown')
    args = parser.parse_args()

    main(args)
