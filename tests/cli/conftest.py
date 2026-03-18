import os
import pytest

CLI = os.getenv("BBCLI", "./bin/bbcli-linux-amd64")


def pytest_collection_modifyitems(items):
    if not os.path.isfile(CLI):
        skip = pytest.mark.skip(reason=f"CLI-Binary nicht gefunden ({CLI}) — zuerst 'make cli' ausführen")
        for item in items:
            item.add_marker(skip)
