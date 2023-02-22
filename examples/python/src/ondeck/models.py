# Code generated by sqlc. DO NOT EDIT.
# versions:
#   sqlc v1.17.1
import dataclasses
import datetime
import enum
from typing import List, Optional


class Status(str, enum.Enum):
    """Venues can be either open or closed"""
    OPEN = "op!en"
    CLOSED = "clo@sed"


@dataclasses.dataclass()
class City:
    slug: str
    name: str


@dataclasses.dataclass()
class Venue:
    """Venues are places where muisc happens"""
    id: int
    status: Status
    statuses: Optional[List[Status]]
    # This value appears in public URLs
    slug: str
    name: str
    city: str
    spotify_playlist: str
    songkick_id: Optional[str]
    tags: Optional[List[str]]
    created_at: datetime.datetime
